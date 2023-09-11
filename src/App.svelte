<script>
import "./app.css"

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
    const response = await fetch('/api/auth/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(userCredentials)
    });
    
    if (response.ok) {
      console.log('response: ', response);
      const userToken = await response.json();
      console.log(userToken);
      // Handle login success, e.g., redirect to homepage
      // window.location.replace('/'
    } else {
      // Handle login failure, e.g., show an error message to the user
      console.log('else ', response);
    }
  } catch (err) {
    console.error(err);
  }
}

</script>

<main>
  <section class="login login__container w-full flex flex-col items-center">
    <div class="login__column column w-5/6 md:4/5 lg:w-1/3">
      <form on:submit={handleLogin} id="login" class="login-form bg-blue-100 flex flex-col px-20 py-2 text-center border border-black">
        <h1 class="login-form__logo">Svelte.IG</h1>
        <input class="login-form__username-input" type="text" id="username" name="username" bind:value={username} required>
        <input class="login-form__email-input" type="text" id="email" name="email" bind:value={email} required>
        <input class="login-form__password-input" type="password" id="password" name="password" bind:value={password} required>
        <input class="login-form__" type="submit" value="Ping">
      </form>
    </div>
  </section>
</main>

<style>

  main {
    height: 100%;
    min-height: 100vh;
    max-height: 100vh;
  }

  .login form input {
    margin-bottom: 6px;
    border: 1px solid #ccc;
  }

</style>
