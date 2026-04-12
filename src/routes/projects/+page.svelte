<script lang="ts">
    import { projects } from "$lib/constants";
    import Project from "../../components/Project.svelte";
    //{[id: string]: string[]}
    let categories: {[id: string] : string[]} = {
        "Roblox" : ["kaboom", "mysteryproj"],
        "Love2D and Other Game Frameworks" : [
            "gooseplatformer3d",
            "playyan",
            "yan",
            "gooserhythm",
            "goosepurgatory",
            "biribiriuofishing",
            "imageplayground",
            "gooseodyssey",
            "gooseplatformer",
            "gooseplatformertogether",
            "gooseinthemines",
            "gooseplatformerendless",
            "goosesurvival",
            "gooseroguelike",
            "spacedefenders"
        ],
        "Rhythm Heaven" : [
            "heavenstudio",
            "rhythmtengokudecomp",
            "tengokumolecano",
            "tengokunightwalkwii",
            "tengokudeltarune",
            "yanimator"
        ],
        "Minecraft" : [
            "treetumblers",
            "nibblesdeepdark",
            "cmbminigames1",
            "cmbminigames2",
        ],
        "Miscallaneous" : [
            "biribiri",
            "bmp2goose",
            "mixtaper",
            "yansstupiddecks",
            "chromapperdraggablewindows",
            "chromappermenubg"
        ]
    }
</script>

<div class="w-full h-110 md:h-150 absolute -z-10">
    <img src="../banner2.png" alt="banner2" class="w-full h-full object-cover object-top-left mask-b-from-70% animate-fadein" style="image-rendering: pixelated; animation-name: flicker; animation-duration: 5s; animation-iteration-count: infinite; animation-delay: 1s;">
</div>
<div class="w-full flex-col pt-80 lg:pt-30">
    <h1 class="text-6xl lg:text-8xl text-center w-full text-shadow-lg/50 animate-fadein">my projects!!</h1> 
</div>

<div class="lg:p-30 p-5 mt-30">
    <button class="w-full pt-5 border-8 border-[#7BE7FF] hover:border-blue-500 transition-all relative">
        <p class="text-center text-6xl">back to main page</p>

        <a class="absolute top-0 bottom-0 left-0 right-0" href="/" title=""></a>
    </button>
    {#each Object.entries(categories) as [name, category]}
        <h1 class="text-4xl lg:text-6xl text-center p-10">{name}</h1>
        {#each category as id}
        <Project name={projects[id].name} image={projects[id].image} description={projects[id].description} footer={projects[id].footer} link={projects[id].link}/>
        {/each}
    {/each}
</div>

<div class="w-full h-110 md:h-150 -z-10">
    <img src="../footer.png" alt="footer" class="w-full h-full object-cover object-top-left mask-t-from-80% animate-fadein" style="image-rendering: pixelated;">
</div>
