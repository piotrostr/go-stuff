import * as components from "./mainComponents"
export * from "./mainComponents"

/**
 * @description 
 */
export function getUser() {
	return webapi.get<components.Response>("/users/id/:userId")
}

/**
 * @description 
 */
export function createUser() {
	return webapi.post<null>("/users/create")
}
