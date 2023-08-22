<script lang="ts">
    import type {Player} from "../types/player";

    export let player:Player
</script>

<div class="flex flex-row">
    <div>{player.Id}</div>
    <div>{player.Name}</div>
</div>