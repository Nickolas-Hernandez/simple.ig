<script>
let username;
let email;
let password;


async function handleLogin(e) {
  e.preventDefault();

  const userCredentials = {
    username,
    email,
    password
  };

  try{
    const response = await fetch('/api/user/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(userCredentials)
    });
    
    if (response.ok) {
      // Handle successful login, e.g., store authentication token in browser's local storage
      console.log('success!');
      const data = await response.json();
      console.log(data);
    } else {
      // Handle login failure, e.g., show an error message to the user
      console.log('fail :(');
    }
  } catch {
    console.log('error');
  }
}

</script>

<main>
  <section class="login">
    <h1>Create an Account</h1>
    <form on:submit={handleLogin} id="login">
      <label for="username">Username</label>
      <input type="text" id="username" name="username" bind:value={username} required>
      <label for="email">Email</label>
      <input type="text" id="email" name="email" bind:value={email} required>
      <label for="password">Password</label>
      <input type="password" id="password" name="password" bind:value={password} required>
      <input type="submit" value="Ping">
    </form>
  </section>
</main>

<style>

</style>
