import {expect, assert} from "chai";
import "mocha"
import axios from "axios";
import connect from "./test_utils/db.js";
import {createTestUser, createUserAndGetToken} from "./test_utils/users.js";

describe('CRUD routes of connections should work fine', () => {
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
		await connection.query('DELETE FROM users WHERE username = $1', ['test2']);
		await connection.query('DELETE FROM connections WHERE common_name = $1', ['test']);
		await connection.query('DELETE FROM connections WHERE common_name = $1', ['test2']);
	});

	it('Should create a connection successfully', async () => {
		let authToken = await createUserAndGetToken(apiUrl);
		const connectionData = {
			"common_name": "test",
			"db_name": "test",
			"host": "127.0.0.1",
			"port": 5432, // Debe ser tipo int
			"username": "postgres",
			"password": "4c9642ea7580a191"
		};

		await axios.post(`${apiUrl}/connections`, connectionData, {
			headers: {
				Authorization: "Bearer " + authToken
			}
		}).then((res) => {
			expect(res.status).to.equal(200);
		}).catch((err) => {
			assert.fail("Create connection route failed");
		});

		await connection.query('DELETE FROM connections WHERE common_name = $1', ['test']);
	});

	it('Should delete a connection successfully', async () => {
		let authToken = await createUserAndGetToken(apiUrl);
		const connectionData = {
			"common_name": "test",
			"db_name": "test",
			"host": "127.0.0.1",
			"port": 5432, // Debe ser tipo int
			"username": "postgres",
			"password": "4c9642ea7580a191"
		};

		let id;
		await axios.post(`${apiUrl}/connections`, connectionData, {
			headers: {
				Authorization: "Bearer " + authToken
			}
		}).then((res) => {
			expect(res.status).to.equal(200);
			id = res.data.id;
		}).catch((err) => {
			assert.fail("Create connection route failed");
		});

		await axios.delete(`${apiUrl}/connections/${id}`, {
			headers: {
				Authorization: "Bearer " + authToken
			}
		}).then((res) => {
			expect(res.status).to.equal(200);
		}).catch((err) => {
			assert.fail("Delete connection route failed");
		});

		await connection.query('DELETE FROM connections WHERE common_name = $1', ['test']);
	});

	it("Should not allow to delete a connection that doesn't exist", async () => {
		let authToken = await createUserAndGetToken(apiUrl);
		await axios.delete(`${apiUrl}/connections/9999999`, {
			headers: {
				Authorization: "Bearer " + authToken
			}
		}).then((res) => {
			assert.fail("Delete connection route failed");
		}).catch((err) => {
			expect(err.response.status).to.equal(404);
		});

	});

	it('Should not allow to delete a connection that is not yours', async () => {
		let authToken = await createUserAndGetToken(apiUrl);
		const connectionData = {
			"common_name": "test",
			"db_name": "test",
			"host": "127.0.0.1",
			"port": 5432, // Debe ser tipo int
			"username": "postgres",
			"password": "4c9642ea7580a191"
		};

		let id;
		await axios.post(`${apiUrl}/connections`, connectionData, {
			headers: {
				Authorization: "Bearer " + authToken
			}
		}).then((res) => {
			expect(res.status).to.equal(200);
			id = res.data.id;
		}).catch((err) => {
			assert.fail("Create connection route failed");
		});

		let randomPassword = Math.random().toString(36).substring(7);
		await axios.post(`${apiUrl}/auth/signup`, {
			email: "test2@gmail.com",
			password: randomPassword,
			name: "test2",
			username: "test2"
		})
		const token = await axios.post(`${apiUrl}/auth/login`, {
			usernameOrEmail: "test2@gmail.com",
			password: randomPassword
		}).then((res) => {
			return res.data.token
		}).catch((err) => {
			return ""
		})

		await axios.delete(`${apiUrl}/connections/${id}`, {
			headers: {
				Authorization: "Bearer " + token
			}
		}).then((res) => {
			assert.fail("Delete connection route failed");
		}).catch((err) => {
			expect(err).to.not.be.null;
		});

	});

});