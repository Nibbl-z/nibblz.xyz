<script lang="ts">
    import { onMount } from "svelte";

    let topbar = [
        ["home", "/"],
        ["my projects", "/projects/"],
        ["my thoughts", "/thoughts/"]
    ]

    let page = $state("/")

    onMount(() => {
        page = window.location.pathname
        console.log(page)
    })
</script>

<div class="w-full flex flex-row items-center justify-center bg-linear-to-b from-[#23172b] to-black">
    {#each topbar as item}
        <div class="pl-5 pr-5 pt-2 pb-2">
            {#if page == item[1]}
                <a href={item[1]} class="text-2xl [text-shadow:0px_0px_12px_#3584e4] text-blue-400! hover:text-blue-500!">{item[0]}</a>
            {:else}
                <a href={item[1]} class="text-2xl hover:[text-shadow:0px_0px_12px_#3584e4] hover:text-blue-300! transition-all">{item[0]}</a>
            {/if}
            
        </div>
    {/each}
</div>