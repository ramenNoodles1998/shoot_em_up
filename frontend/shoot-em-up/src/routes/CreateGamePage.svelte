<script>
    import { onMount } from 'svelte';
    let conn;
    let msg;
    let messages = ['hello'];

    const onSubmit = () => {
        if (!conn) {
            return false;
        }
        if (!msg) {
            return false;
        }
        conn.send(msg);
        msg = '';
        return false;
    }

    onMount(() => {
        conn = new WebSocket('ws://127.0.0.1:8080/ws');
        conn.onclose = function (evt) {
            messages.push('Connection closed')
        };
        conn.onmessage = function (evt) {
            let eventMessages = evt.data.split('\n');
            for (let message of eventMessages) {
                messages.push(message)
            }
            console.log(messages)
        };
    });
</script>

<div>
    <div id="log">
        <ul>
            <li>list</li>
            {#each messages as message}
                <li>{message}</li>
            {/each} 
        </ul>
    </div>
    <div id="form">
        <button on:click={onSubmit}>Send</button>
        <input bind:value={msg} on:submit={onSubmit} type="text" size="64"/>
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
}

#form {
    padding: 0 0.5em 0 0.5em;
    margin: 0;
    position: absolute;
    bottom: 1em;
    left: 0px;
    width: 100%;
    overflow: hidden;
}

</style>