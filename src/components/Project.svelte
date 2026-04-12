<script lang="ts">
    import { fade } from 'svelte/transition';
    import { onMount } from "svelte";

    let {name, image, description, footer = undefined, link}: {name: string, image: string, description: string, footer: string | undefined, link: string} = $props()

    let inView = $state(false)
    let element: Element;

    onMount(() => {
        const observer = new IntersectionObserver((entries) => {
            entries.forEach((entry) => {
                inView = entry.isIntersecting
            });
        });

        observer.observe(element);
    })
</script>


<div bind:this={element}>
    <div class="w-full border-8 border-[#0821FF] backdrop-blur-lg mb-10" class:notvisible={!inView} class:fade>
        <div class="w-full h-full absolute outline-8 outline-[#0821FF] blur-lg -z-10"></div>
        <div class="p-5">
            <h1 class="text-4xl lg:text-6xl">{name}</h1>
            <div class="flex xl:flex-row flex-col pt-5">
                <p class="text-2xl lg:text-3xl xl:w-[50%] p-2">{description}</p>
                {#if !image.includes("youtube.com")}
                    <img src={image} alt={name} class="xl:w-[50%] max-h-130 object-contain p-2 rounded-2xl"/>
                {:else}
                    <iframe class="xl:w-[50%] aspect-video object-contain p-2 rounded-2xl" src={image} title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
                {/if}
                
            </div>
            {#if footer}
                <p class="text-xl text-center text-gray-500!">{footer}</p>
            {/if}
            <button class="w-full pt-5 border-8 border-[#7BE7FF] hover:border-blue-500 transition-all relative">
                <p class="text-center text-6xl">Check it out!</p>
                <a class="absolute top-0 bottom-0 left-0 right-0" href={link} target="_blank" title=""></a>
            </button>
        </div>
        
    </div>
</div>
