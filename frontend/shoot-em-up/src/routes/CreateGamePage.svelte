<script>
    import { onMount } from 'svelte';
    let conn;
    let health = 500;
    let opponentHealth = 500;
    let healthpx = '500px';
    let opponentHealthpx = '500px';
    let display = '';
    let firstMessage = true;
    let id;

    const onAttack = () => {
        if (!conn) {
            return false;
        }
        conn.send(JSON.stringify({id, damage: 100}));
        return false;
    }

    onMount(() => {
        conn = new WebSocket('ws://127.0.0.1:8080/ws');
        conn.onclose = function (evt) {
            health = 0;
        };

        conn.onmessage = function (evt) {
            let eventMessages = evt.data.split('\n');
            if(firstMessage) {
                firstMessage = false;
                id = eventMessages[0];
            } else {
                for (let message of eventMessages) {
                    console.log(message)
                    let data = JSON.parse(message);
                    if(data.id === id) continue;
                    if(data.damage) {
                        if(health > 0) {
                            health -= parseInt(data.damage);
                            conn.send(JSON.stringify({id, hit: data.damage}));
                        } else {
                            display = 'You Dead';
                            conn.send(JSON.stringify({id, hit: data.damage}));
                        }
                    }

                    if(data.hit) {
                        console.log(opponentHealth);
                        if(opponentHealth > 0) {
                            opponentHealth -= parseInt(data.hit);
                        } else {
                            display = 'You killed the enemy';
                        }
                    }
                }
                healthpx = health + 'px';
                opponentHealthpx = opponentHealth + 'px';
            }
        };
    });
</script>

<div>
    <div id="log">
        <h1 style:color="red">
            {display}
        </h1>
        <p class="health" style:height={healthpx}>Your HP</p>
        <p class="opponent" style:height={opponentHealthpx}>Enemy HP</p>
    </div>
    <div id="buttons">
        <button on:click={onAttack}>Attack Enemy</button>
    </div>
</div>

<style>
#log {
    background: white;
    color: black;
    margin: 0;
    padding: 0.5em 0.5em 0.5em 0.5em;
    position: absolute;
    top: 0.5em;
    left: 0.5em;
    right: 0.5em;
    bottom: 3em;
    overflow: auto;
    border-radius: .5em;
}

.health {
    color: white;
    background-color: rgb(141, 67, 67);
    position: absolute;
    top: auto;
    left: 2em;
    right: auto;
    bottom: 2em;
    padding: 2em;
}

.opponent {
    color: white;
    background-color: rgb(141, 67, 67);
    position: absolute;
    top: auto;
    left: auto;
    right: 2em;
    bottom: 2em;
    padding: 2em;
}
#buttons {
    display: flex;
    justify-content: flex-start;
    padding: 0.5em;
    margin: 0;
    background-color: rgb(141, 67, 67);
    position: absolute;
    bottom: 1em;
    left: 0.5em;
    right: .5em;
    overflow: hidden;
    border-bottom-right-radius: .5em;
    border-bottom-left-radius: .5em;
}

</style>