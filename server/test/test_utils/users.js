import axios from "axios";

export async function createTestUser(password, apiUrl) {
	await axios.post(`${apiUrl}/auth/signup`, {
		email: "test@gmail.com",
		password: password,
		name: "test",
		username: "test"
	})
}

export async function createUserAndGetToken(apiUrl) {
	let randomPassword = Math.random().toString(36).substring(7);
	await createTestUser(randomPassword, apiUrl)
	const token = await axios.post(`${apiUrl}/auth/login`, {
		usernameOrEmail: "test@gmail.com",
		password: randomPassword
	}).then((res) => {
		return res.data.token
	}).catch((err) => {
		return ""
	})
	return token
}