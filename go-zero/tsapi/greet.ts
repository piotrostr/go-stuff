import * as components from "./greetComponents"
export * from "./greetComponents"

/**
 * @description 
 * @param params
 */
export function greet(params: components.RequestParams) {
	return webapi.get<components.Response>("/greet/from/:name", params)
}
