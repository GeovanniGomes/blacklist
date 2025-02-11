package queue

// import (
// 	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
// )

// type ConsumerLoader struct {
// 	brokenHandler contracts.IQueue
// }

// func NewConsumerLoader(brokenHandler contracts.IQueue) *ConsumerLoader {

// 	return &ConsumerLoader{brokenHandler: brokenHandler}
// }

// // func (c *ConsumerLoader) StartConsumers(container depedence_injector.ContainerInjection) {
// // 	blacklistNotifyConsumer := consumer.NewBlacklistConsumer(c.brokenHandler)

// // 	consumers := map[string]func([]byte) error{
// // 		os.Getenv("QUEUE_BLACKLIST"):        blacklistNotifyConsumer.HandleMessage(),
// // 		os.Getenv("QUEUE_REPORT_BLACKLIST"): blacklistConsumerReport.HandleMessage(),
// // 	}

// for queueName, handler := range consumers {
// 	go func(queue string, h func([]byte) error) {
// 		err := c.brokenHandler.Consume(queue, func(msg []byte) {
// 			if err := h(msg); err != nil {
// 				log.Printf("❌ Erro ao processar mensagem da fila %s: %v", queue, err)
// 			}
// 		})
// 		if err != nil {
// 			log.Printf("❌ Erro ao iniciar consumidor %s: %v", queue, err)
// 		}
// 	}(queueName, handler)
// }

// // 	log.Println("🎯 Todos os consumidores foram iniciados.")
// // }
