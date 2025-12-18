<script lang="ts">
    export let name: string
    export let image: string
    export let description: string
    export let footer: string | undefined = undefined
    export let link: string
</script>

<div class="w-full border-8 border-[#0821FF] backdrop-blur-lg mb-10">
    <div class="w-full h-full absolute outline-8 outline-[#0821FF] blur-lg -z-10"></div>
    <div class="p-5">
        <h1 class="text-6xl">{name}</h1>
        <div class="flex xl:flex-row flex-col pt-5">
            <p class="text-3xl xl:w-[50%] p-2">{description}</p>
            {#if !image.includes("youtube.com")}
                <img src={image} alt={name} class="xl:w-[50%] max-h-130 object-contain p-2 rounded-2xl"/>
            {:else}
                <iframe class="xl:w-[50%] aspect-video object-contain p-2 rounded-2xl" src={image} title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
            {/if}
            
        </div>
        {#if footer}
            <p class="text-xl text-center text-gray-500!">{footer}</p>
        {/if}
        <button class="w-full pt-5 border-8 border-[#7BE7FF] hover:border-[#0821FF]">
            <a class="text-center text-6xl" href={link} target="_blank">Check it out!</a>
        </button>
    </div>
    
</div>