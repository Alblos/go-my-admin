import {expect, assert} from "chai";
import "mocha"
import axios from "axios";
import connect from "./test_utils/db.js";
import {createTestUser} from "./test_utils/users.js";

describe('Authorization routes should work fine', () => {
	const apiUrl = 'http://localhost:3000';
	let connection;

	before(async () => {
		connection = await connect();
	});

	after(async () => {
		await connection.end();
	});

	afterEach(async () => {
		await connection.query('DELETE FROM users WHERE username = $1', ['test']);
	});

	it('Signup route should work perfectly', async () => {
		await axios.post(`${apiUrl}/auth/signup`, {
			email: "test@gmail.com",
			password: "123456",
			name: "test",
			username: "test"
		}).then((res) => {
			expect(res.status).to.equal(200);
		}).catch((err) => {
			assert.fail("Signup route failed");
		});
	});

	it('Login route should work perfectly', async () => {
		let randomPassword = Math.random().toString(36).substring(7);
		await createTestUser(randomPassword, apiUrl)
		await axios.post(`${apiUrl}/auth/login`, {
			usernameOrEmail: "test@gmail.com",
			password: randomPassword
		}).then((res) => {
			expect(res.status).to.equal(200);
			expect(res.data).to.have.property('token');
			expect(res.data.token).to.be.a('string');
		}).catch((err) => {
			assert.fail("Login route failed");
		});
	});

	it('Should not login if the password is not correct', async() => {
		let randomPassword = Math.random().toString(36).substring(7);
		await createTestUser(randomPassword, apiUrl)
		await axios.post(`${apiUrl}/auth/login`, {
			usernameOrEmail: "test@gmail.com",
			password: "wrongPassword"
		}).then((res) => {
			assert.fail("Should not login if the password is not correct");
		}).catch((err) => {
			expect(err.response.status).to.equal(401);
			expect(err.response.data).to.have.property('error');
			expect(err.response.data.error).to.be.a('boolean');
			expect(err.response.data.error).to.equal(true);
		});
	});
});