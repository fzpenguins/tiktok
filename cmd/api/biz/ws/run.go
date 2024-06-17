package ws

import (
	"context"
	"encoding/json"
	"log"
	"tiktok/cmd/api/dal/cache"
	"tiktok/cmd/api/dal/db"
	"tiktok/cmd/api/dal/db/dao"
	"tiktok/cmd/api/dal/mq"
	"tiktok/pkg/constants"
	"time"
)

func MessageHandler() {
	for {
		select {
		case broadcast := <-Broadcast:
			//log.Println("建立新连接")
			log.Println("successfully come into broadcast goroutine")
			msgDao := dao.GetMsgDao(context.Background())

			if broadcast.Type == constants.SingleChat { //单聊
				msgJSON, _ := json.Marshal(broadcast)

				//存储到数据库

				if !Clients[ClientsSingleMap[broadcast.FromUid]] { //对方有无上线
					err := mq.PublishMsg(msgJSON, broadcast.Target)
					if err != nil {
						log.Println(err)
						continue
					}
					log.Println("对方未上线")
					err = msgDao.StoreSingleChatMsg(broadcast.FromUid, broadcast.Target, broadcast.Content, false)
					if err != nil {
						log.Println(err)
						continue
					}
					continue
				}
				select {
				case ClientsSingleMap[broadcast.FromUid].send <- msgJSON: //发送
					err := msgDao.StoreSingleChatMsg(broadcast.FromUid, broadcast.Target, broadcast.Content, true)
					if err != nil {
						log.Println(err)
						continue
					}
				default:
					close(ClientsSingleMap[broadcast.FromUid].send)

					delete(ClientsSingleMap, broadcast.FromUid)
				}

			} else if broadcast.Type == constants.GroupChat { //群聊
				var err error
				//var members = make([]string, e.MaxStore)
				var members []string
				members, err = cache.RedisClient.SMembers(context.Background(), broadcast.Target).Result()
				if err != nil {
					log.Println(err)
				}
				msgJSON, _ := json.Marshal(broadcast)
				onlineMember := ClientsMap[broadcast.Target] //这里的问题只存入一个，为什么会这样
				for _, member := range members {

					if member == broadcast.FromUid {
						continue
					}

					msg := &db.Message{
						CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
						DeletedAt: "",
						FromUid:   member,
						ToUid:     broadcast.Target,
						Type:      constants.GroupChat,
						Content:   broadcast.Content,
					}

					if conn, ok := onlineMember[member]; ok {
						conn.send <- msgJSON
						msg.ReadTag = true
					} else {
						//log.Println(member)
						msg.ReadTag = false
					}

					if err := dao.NewDBClient(context.Background()).Create(msg); err != nil {
						log.Println(err)
						continue
					}
				}
			}
		case conn := <-Unregister:
			log.Println("结束连接")
			ClientsLock.Lock()

			if conn.Type == constants.SingleChat {
				if Clients[conn] {
					delete(ClientsSingleMap, conn.FromUid)
					delete(Clients, conn)
					close(conn.send)
				}
			} else if conn.Type == constants.GroupChat {
				delete(Clients, conn)
				delete(ClientsMap[conn.Target], conn.FromUid)
				close(conn.send)
			}
			if len(ClientsMap[conn.Target]) == 0 {
				delete(ClientsMap, conn.Target)
			}

			ClientsLock.Unlock()
			log.Println("已完全结束连接")
		case conn := <-Register:
			log.Println("进行register相关操作")

			if conn.Type == constants.GroupChat {
				log.Println(conn.Target)

				if _, ok := ClientsMap[conn.Target]; !ok {
					ClientsMap[conn.Target] = make(map[string]*Client) //不能反复make，有make的就不能再次make了
				}

				ClientsLock.Lock()
				Clients[conn] = true
				ClientsMap[conn.Target][conn.FromUid] = conn
				ClientsLock.Unlock()

				cache.RedisClient.SAdd(context.Background(), conn.Target, conn.FromUid)
			} else if conn.Type == constants.SingleChat {

				ClientsLock.Lock()
				Clients[conn] = true
				ClientsSingleMap[conn.Target] = conn //目标映射Client
				ClientsLock.Unlock()

			}

		}
	}
}
