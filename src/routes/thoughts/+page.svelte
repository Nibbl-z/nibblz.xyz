<script lang="ts">
    import { marked, Renderer, type Tokens } from "marked";
    import Post from "../../components/Post.svelte";
    import Topbar from "../../components/Topbar.svelte";
    import { fade } from 'svelte/transition';
    import { Tween } from "svelte/motion";
    import { linear } from "svelte/easing";

    const renderer = new Renderer()

    renderer.heading = (token: Tokens.Heading) => {
        const {text, depth} = token
        const levels: Record<number, string> = {
            1: "text-5xl",
            2: "text-4xl",
            3: "text-3xl"
        }

        return `<h${depth} class="${levels[depth]}">${text}</h${depth}>`
    }

    const opacity = new Tween(0, {
        duration: 200,
        easing: linear
    })

    renderer.list = (token: Tokens.List) => {
        const items = token.items.map(item => {
            return `<li>${item.text}</li>`
        }).join("\n")

        if (token.ordered) {
            return `<ol class="list-decimal pl-5">${items}</ol>`
        } else {
            return `<ul class="list-disc pl-5">${items}</ul>`
        }
    }

    marked.use({renderer})

    let postOpen = $state(false)

    function openPost() {
        postOpen = true
        opacity.set(1.0)
    }

    function closePost() {
        postOpen = false
        opacity.set(0.0)
    }
</script>

<div class="fixed w-screen h-screen z-9 p-0 backdrop-blur-sm" style="opacity: {opacity.current}; pointer-events: {postOpen ? "auto" : "none"}">
    <button class="w-[10em] border-8 m-3 border-[#7BE7FF] hover:border-blue-500 transition-all absolute z-10" onclick={closePost}>
        <div class="w-full h-full absolute bg-black -z-9 opacity-70"></div>
        <div class="p-2">
            <p class="text-center text-6xl">back</p>
        </div>
    </button>

    <div class="w-full h-full md:p-30 p-5 pt-30">
        <div class="border-8 border-[#0821FF] backdrop-blur-lg w-full h-full ">
            <div class="w-full h-full absolute outline-8 outline-[#0821FF] blur-lg -z-10"></div>
            <div class="w-full h-full absolute bg-black -z-9 opacity-70"></div>
            <div class="p-5">
                <p class="text-6xl text-center">title</p>
                <div>
                    {@html marked.parse("# helloo \n ## hello! \n ### hellooooo !!!! \n **hai!!!** and *hai* \n- big things are happening \n- reallly big things")}
                </div>
            </div>
        </div>
    </div>
</div>

<Topbar/>

<div class="w-full h-110 md:h-150 absolute -z-10">
    <img src="../banner2.png" alt="banner2" class="w-full h-full object-cover object-top-left mask-b-from-70% animate-fadein" style="image-rendering: pixelated;">
</div>
<div class="w-full flex-col pt-80 lg:pt-30">
    <h1 class="text-6xl lg:text-8xl text-center w-full text-shadow-lg/50 animate-fadein">my thoughts!!</h1> 
</div>

<div class="pt-30">
    <p class="text-4xl text-center pl-3 pr-3 opacity-0" style="animation: fadein 1s 1 0.1s; animation-fill-mode: forwards">a place to share any thing that i think...</p>
</div>

<div class="flex flex-row items-center flex-wrap justify-center pl-5 pr-5">
    <Post name="post name" description="bla bla blaaaa" image="/biribiriuo.webp" timestamp={1780534756} click={openPost}/>
    <Post name="post name" description="bla bla blaaaa" image="/biribiriuo.webp" timestamp={1780534756} click={openPost}/>
    <Post name="post name" description="bla bla blaaaa" image="/biribiriuo.webp" timestamp={1780534756} click={openPost}/>
    <Post name="post name" description="bla bla blaaaa" image="/biribiriuo.webp" timestamp={1780534756} click={openPost}/>
</div>