<script>
  import { Router, Route, Link } from 'svelte-routing';
  import CreateGamePage from './routes/CreateGamePage.svelte';

  export let url = '';
</script>

<Router url="{url}">
  <div class="home">
    <nav class="buttons">
        <Link to="/create-game">Create Game</Link> 
    </nav>
    <Route path="create-game" component="{CreateGamePage}"></Route> 
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
