<script lang='ts'>
	let proceedDisabled = true;
	let name: string;

	async function doRegister() {
		const res = await fetch('http://localhost:8000/register', {
			method: 'POST',
			body: JSON.stringify({
				name
			})
		});

		const availablePlayers = await res.json();

		localStorage.setItem('availablePlayers', JSON.stringify(availablePlayers));
		proceedDisabled = false;
	}
</script>

<h1>Register</h1>
<label for='name'>Enter your name</label>
<input id='name' type='text' bind:value={name}>
<button on:click={doRegister}>Submit</button>

<p><a href='/game' hidden={proceedDisabled}>Proceed</a></p>