<script lang="ts">
    import {onMount} from "svelte";
    import type {Player} from "../../lib/types/player";
    import PlayerDisplay from "../../lib/componets/PlayerDisplay.svelte";

    let playerData:Player[] = []
    onMount(async () => {
        let response = await fetch("/api/players");
        playerData = await response.json()
    })
</script>
<div>
    <h1>players</h1>
    {#each playerData as player}
        <PlayerDisplay player={player}></PlayerDisplay>
    {/each}
</div>