<script lang="ts">
    import { projects } from "$lib/constants";
    import Project from "../components/Project.svelte";
    
    let langs = [
        ["https://raw.githubusercontent.com/vscode-icons/vscode-icons/c15969c0d34ac738bde5a6416fc9c79727eaee37/icons/file_type_luau.svg", "Luau"],
        ["https://raw.githubusercontent.com/vscode-icons/vscode-icons/c15969c0d34ac738bde5a6416fc9c79727eaee37/icons/file_type_rust.svg", "Rust"],
        ["https://raw.githubusercontent.com/vscode-icons/vscode-icons/c15969c0d34ac738bde5a6416fc9c79727eaee37/icons/file_type_lua.svg", "Lua"],
        ["https://raw.githubusercontent.com/vscode-icons/vscode-icons/b07f49907d1da36efabaa5a18989cb66aef006d3/icons/file_type_go_gopher.svg", "Go"],
        ["https://raw.githubusercontent.com/vscode-icons/vscode-icons/b07f49907d1da36efabaa5a18989cb66aef006d3/icons/file_type_svelte.svg", "Svelte"],
        ["https://raw.githubusercontent.com/vscode-icons/vscode-icons/b07f49907d1da36efabaa5a18989cb66aef006d3/icons/file_type_typescript_official.svg", "Typescript"],
        ["https://raw.githubusercontent.com/vscode-icons/vscode-icons/b07f49907d1da36efabaa5a18989cb66aef006d3/icons/file_type_c.svg", "C"],
        ["https://raw.githubusercontent.com/vscode-icons/vscode-icons/b07f49907d1da36efabaa5a18989cb66aef006d3/icons/file_type_assembly.svg", "Assembly"],
        ["https://raw.githubusercontent.com/vscode-icons/vscode-icons/b07f49907d1da36efabaa5a18989cb66aef006d3/icons/file_type_python.svg", "Python"],
        ["https://raw.githubusercontent.com/vscode-icons/vscode-icons/b07f49907d1da36efabaa5a18989cb66aef006d3/icons/file_type_css.svg", "CSS (I hate.)"]
    ]

    let tools = [
        ["https://upload.wikimedia.org/wikipedia/commons/9/9a/Visual_Studio_Code_1.35_icon.svg", "VS Code"],
        ["https://dl.flathub.org/media/org/love2d/love2d/8230b18f40cc60f448a51e43910c64ce/icons/128x128@2/org.love2d.love2d.png", "Love2D"],
        ["https://upload.wikimedia.org/wikipedia/commons/4/42/Roblox_Studio_Logo_2022.png", "Roblox Studio"],
        ["https://elttob.uk/Fusion/0.3/assets/logo-dark.svg", "Fusion"],
        ["https://avatars.githubusercontent.com/u/4350249?s=200&v=4", "Spigot"],
        ["https://avatars.githubusercontent.com/u/7608950?s=200&v=4", "PaperMC"],
        ["DecompMe.png", "decomp.me"],
        ["https://appwrite.io/assets/logomark/logo.png", "Appwrite"],
        ["https://raw.githubusercontent.com/emilk/egui/refs/heads/main/web_demo/favicon.ico", "egui"]
    ]

    let accounts = [
        ["fa-brands fa-github", "GitHub", "https://github.com/Nibbl-z"],
        ["fa-brands fa-youtube", "YouTube", "https://www.youtube.com/@nibbl_z"],
        ["fa-brands fa-itch-io", "itch.io", "https://nibbl-z.itch.io/"],
        ["newgrounds.png", "Newgrounds", "https://nibbl-z.newgrounds.com"],
        ["fa-brands fa-soundcloud", "Soundcloud", "https://soundcloud.com/nibbl_z"],
        ["roblox.png", "Roblox", "https://www.roblox.com/users/281430862/profile"]
    ]

    let featured_projects = [projects["kaboom"], projects["playyan"], projects["mysteryproj"]]
</script>

<div class="flex flex-col items-center">
    
    <div class="w-full h-110 md:h-150 absolute -z-10">
        <img src="banner.png" alt="banner" class="w-full h-full object-cover object-top-left mask-b-from-70% animate-fadein" style="image-rendering: pixelated;">
    </div>
    <div class="w-full flex-col pt-80 lg:pt-30">
        <h1 class="text-9xl text-center w-full text-shadow-lg/50 animate-fadein">hai, i'm</h1>
        <div>
            <img src="name_plate.png" alt="nibbles" class="h-[15em] object-center w-full object-contain absolute z-10" style="image-rendering: pixelated; animation-name: flicker; animation-duration: 5s; animation-iteration-count: infinite; animation-delay: 1s;">
            <img src="name_plate_blur.png" alt="nibbles" class="h-[15em] object-center w-full object-contain blur-md animate-fadein" style="image-rendering: pixelated">
        </div>
        
    </div>
    <p class="text-4xl pt-5 text-center pl-3 pr-3">i'm a programmer/game developer/music composer who's been making stuff for over 5 years!!</p>
    <p class="text-3xl pt-5 text-center pl-3 pr-3">here are some languages i know sorted by how much i like them!</p>

    <div class="flex flex-row items-center flex-wrap justify-center pl-3 pr-3">
        {#each langs as lang}
        <div>
            <img src={lang[0]} class="w-[6em] h-[6em] p-2 absolute z-10" alt={lang[1]}/>
            <img src={lang[0]} class="w-[6em] h-[6em] p-2 blur-sm opacity-75" alt={lang[1]}/>
            <p class="text-center text-sm">{lang[1]}</p>
        </div>
        
        {/each}
    </div>

    <p class="text-3xl pt-5 text-center pl-3 pr-3">and the same thing but with tools/frameworks/software in no particular order (i love them all)!</p>

    <div class="flex flex-row items-center flex-wrap justify-center pl-3 pr-3">
        {#each tools as tool}
        <div>
            <img src={tool[0]} class="w-[6em] h-[6em] p-2 absolute z-10 object-contain" alt={tool[1]}/>
            <img src={tool[0]} class="w-[6em] h-[6em] p-2 blur-sm object-contain opacity-75" alt={tool[1]}/>
            <p class="text-center text-sm">{tool[1]}</p>
        </div>
        
        {/each}
    </div>

    <p class="text-4xl pt-5 text-center pl-3 pr-3 mt-10">here are my current coolest projects:</p>

    <div class="flex flex-col items-center lg:pl-30 lg:pr-30 pl-3 pr-5">
        <button class="w-full pt-5 border-8 m-10 border-[#7BE7FF] hover:border-[#0821FF]">
            <a class="text-center text-6xl" href="/projects">see all projects!!</a>
        </button>
        {#each featured_projects as project}
            <Project name={project.name} image={project.image} description={project.description} footer={project.footer} link={project.link}/>
        {/each}
    </div>

    <p class="text-4xl pt-5 text-center pl-3 pr-3 mt-10">also my accounts on a buncha stuff:</p>

    <div class="flex flex-row items-center flex-wrap justify-center pl-3 pr-3 mb-10">
        {#each accounts as account}
        <div class="max-w-[4em] max-h-[4em] m-4 flex flex-col items-center">
            <a class="text-[4em]/0" title="" target="_blank" href={account[2]}>
                {#if account[0].includes(".")}
                    <img src={account[0]} alt="">
                {:else}
                    <i class={account[0]}></i>
                {/if}
            </a>
            <p class="text-center text-md">{account[1]}</p>
        </div>
        
        {/each}
    </div>

    <div class="w-full h-110 md:h-150 -z-10">
        <img src="footer.png" alt="footer" class="w-full h-full object-cover object-top-left mask-t-from-80% animate-fadein" style="image-rendering: pixelated;">
    </div>
</div>