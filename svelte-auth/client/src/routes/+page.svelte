<script lang="ts">
	import axios from 'axios';

	// data bindign for ui handling
	let emailBind: string = '';
	let passwordBind: string = '';
	const onLogin = async (e: Event) => {
		e.preventDefault();
		const form = e.target as HTMLFormElement;
		const formData = new FormData(form);

		const username = formData.get('username');
		const password = formData.get('password');

		// form validation
		if (!username || !password) {
			setTimeout(() => {
				alert('Please fill all the fields');
				emailBind = '';
				passwordBind = '';
			}, 1000);
			return;
		}
		try {
			const response = await axios.post(
				'http://localhost:3000/auth',
				{
					username,
					password
				},
				{
					headers: {
						'Content-Type': 'application/x-www-form-urlencoded'
					}
				}
			);
			if (response.status === 200) {
				// deconstruct
				const { data } = response.data;
				const { username, password } = data;
				console.log(username);
				console.log(password);
			}
		} catch (error) {
			if (axios.isAxiosError(error)) {
				console.error(error);
			}
		}
	};

	const fetchServer = async () => {
		try {
			const response = await axios.get('http://localhost:3001/');
			console.log(response.data);
		} catch (err) {
			if (err instanceof Error) {
				console.error(err.message);
			}
		}
	};
</script>

<main class="flex justify-center items-center h-screen overflow-hidden">
	<form action="" on:submit={onLogin} class="flex flex-col items-center">
		<input type="text" name="username" placeholder="Username" bind:value={emailBind} />
		<input type="password" name="password" placeholder="Password" bind:value={passwordBind} />
		<button>Login</button>
	</form>
	<button on:click={fetchServer}>Fetch to Server</button>
</main>
