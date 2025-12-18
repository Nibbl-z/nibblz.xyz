<script lang="ts">
    import { projects } from "$lib/constants";
    import Project from "../../components/Project.svelte";
    //{[id: string]: string[]}
    let categories: {[id: string] : string[]} = {
        "Roblox" : ["kaboom", "mysteryproj"],
        "Non-Roblox Games" : [
            "playyan",
            "gooserhythm",
            "goosepurgatory",
            "biribiriuofishing",
            "imageplayground",
            "yan",
            "gooseodyssey",
            "gooseplatformer",
            "gooseplatformertogether",
            "gooseplatformerendless",
            "goosesurvival",
            "gooseroguelike",
            "spacedefenders"
        ],
        "Rhythm Heaven" : [
            "heavenstudio",
            "rhythmtengokudecomp",
            "tengokunightwalkwii",
            "tengokudeltarune",
            "yanimator"
        ]
    }
</script>

<h1 class="text-9xl text-center p-10">projectss</h1>

<div class="p-30">
    {#each Object.entries(categories) as [name, category]}
        <h1 class="text-6xl text-center p-10">{name}</h1>
        {#each category as id}
        <Project name={projects[id].name} image={projects[id].image} description={projects[id].description} footer={projects[id].footer} link={projects[id].link}/>
        {/each}
    {/each}
</div>
