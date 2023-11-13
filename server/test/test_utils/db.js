import pkg from 'pg';
const {Client} = pkg;
export default async function connect() {
	const client = new Client({
		host: "localhost",
		user: "postgres",
		password: "562bb3889e26d8fe",
		database: "gomyadmin",
		port: 5432
	});

	await client.connect();
	return client;
}