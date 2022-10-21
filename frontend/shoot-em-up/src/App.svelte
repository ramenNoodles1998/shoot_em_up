<script>
  import { Router, Route, Link } from 'svelte-routing';
  import GamePage from './routes/GamePage.svelte';

  export let url = '';
  let code = 0

  const goToGamePage = () => {
    window.location.href = `http://127.0.0.1:5173/create-game?code=${code}`;
  };
</script>

<Router url="{url}">
  <div class="home">
    <h1 class="title">SHOOT EM UP</h1>
    <nav class="buttons">
        <Link to="/create-game">Create Game</Link> 
    </nav>
    <input placeholder="Enter code to join room" bind:value={code}/>
    <button style:background-color="blue" on:click={goToGamePage}>join</button>
    <Route path="create-game" component="{GamePage}"></Route> 
  </div>
</Router>

<style>
  .title {
    padding-top: 50px;
    margin: auto 0;
    border-radius: 50px;
  }

  .home {
    margin: auto 0;
    border-radius: 50px;
    background-color: rgb(22, 21, 21);
    width: 750px;
    height: 800px;
  }

  .buttons {
    padding-top: 50px;
  }
</style>
