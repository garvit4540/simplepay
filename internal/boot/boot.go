package boot

import (
	"fmt"
	"github.com/garvit4540/simplepay/internal/database"
	"github.com/garvit4540/simplepay/internal/keys"
	"github.com/garvit4540/simplepay/internal/merchants"
	"github.com/garvit4540/simplepay/internal/orders"
	"github.com/garvit4540/simplepay/internal/payments"
	"github.com/garvit4540/simplepay/internal/providerfactory"
	"github.com/garvit4540/simplepay/internal/registry"
	"github.com/garvit4540/simplepay/internal/terminals"
	_ "github.com/lib/pq"
	"log"
)

// Initialize initializes the application
func Initialize() error {
	if err := database.InitializeDatabase(); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	if err := database.RunMigrations(); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	sqlDb, err := database.DatabaseClient.DB()
	if err != nil {
		return fmt.Errorf("failed to convert gorm client to sql client with err : %w", err)
	}

	// Register Services
	registry.InitialiseServiceRegister()

	providerClient := providerfactory.NewProviderService(providerfactory.NewProviderRepo(sqlDb))
	registry.RegisterService(registry.ProviderService, providerClient)

	merchantClient := merchants.NewMerchantService(merchants.NewMerchantRepo(sqlDb))
	registry.RegisterService(registry.MerchantService, merchantClient)

	keysClient := keys.NewKeysService(keys.NewKeysRepo(sqlDb))
	registry.RegisterService(registry.KeysService, keysClient)

	terminalsClient := terminals.NewTerminalService(terminals.NewTerminalRepo(sqlDb))
	registry.RegisterService(registry.TerminalService, terminalsClient)

	ordersClient := orders.NewOrdersService(orders.NewOrdersRepo(sqlDb))
	registry.RegisterService(registry.OrdersService, ordersClient)

	paymentsClient := payments.NewPaymentsService(payments.NewPaymentsRepo(sqlDb))
	registry.RegisterService(registry.PaymentsService, paymentsClient)

	log.Println("Application initialized successfully")
	return nil
}

// Cleanup performs cleanup operations
func Cleanup() error {
	sqlDB, err := database.DatabaseClient.DB()
	if err != nil {
		return fmt.Errorf("failed to convert gorm db to sql db client with err : %w", err)
	}

	err = sqlDB.Close()
	if err != nil {
		return fmt.Errorf("failed to close sql db connection with err : %w", err)
	}

	return nil
}
