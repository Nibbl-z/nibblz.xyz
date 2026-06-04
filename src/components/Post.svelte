<script lang="ts">
    import { fade } from 'svelte/transition';
    import { onMount } from "svelte";

    let {name, image, description, timestamp, click}: {
        name: string, 
        image: string, 
        description: string, 
        timestamp: number,
        click: Function
    } = $props()

    let inView = $state(false)
    let element: Element;

    onMount(() => {
        const observer = new IntersectionObserver((entries) => {
            entries.forEach((entry) => {
                inView = entry.isIntersecting
            });
        });

        const date = new Date(timestamp * 1000)
        console.log(date.toLocaleString())

        observer.observe(element);
    })
</script>


<div bind:this={element}>
    <div class="border-8 border-[#0821FF] backdrop-blur-lg m-4" class:notvisible={!inView} class:fade>
        <div class="w-full h-full absolute outline-8 outline-[#0821FF] blur-lg -z-10"></div>
        <div class="p-5">
            <h1 class="text-4xl lg:text-6xl text-center">{name}</h1>
            <div class="flex xl:flex-row flex-col pt-5">
                <img src={image} alt={name} class="xl:w-[50%] max-h-50 object-contain p-2 rounded-2xl"/>  
                <p class="text-2xl lg:text-3xl xl:w-[50%] p-2">{description}</p>
            </div>
            <p class="text-xl text-center text-gray-500!">{new Date(timestamp * 1000).toLocaleString()}</p>
            <button class="w-full pt-5 border-8 border-[#7BE7FF] hover:border-blue-500 transition-all relative" onclick={() => click()}>
                <p class="text-center text-6xl">read!</p>
            </button>
        </div>
    </div>
</div>