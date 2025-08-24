package registry

const ProviderService string = "provider_service"
const MerchantService string = "merchant_service"
const KeysService string = "keys_service"
const TerminalService string = "terminal_service"
const OrdersService string = "orders_service"

var serviceRegister map[string]interface{}

func InitialiseServiceRegister() {
	serviceRegister = make(map[string]interface{}, 0)
}

func RegisterService(service string, client interface{}) {
	serviceRegister[service] = client
}

func GetServiceFromRegister(service string) interface{} {
	return serviceRegister[service]
}
