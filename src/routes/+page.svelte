<script lang="ts">
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
        ["https://love2d.org/favicon.ico", "Love2D"],
        ["https://upload.wikimedia.org/wikipedia/commons/4/42/Roblox_Studio_Logo_2022.png", "Roblox Studio"],
        ["https://elttob.uk/Fusion/0.3/assets/logo-dark.svg", "Fusion"],
        ["https://avatars.githubusercontent.com/u/4350249?s=200&v=4", "Spigot"],
        ["https://avatars.githubusercontent.com/u/7608950?s=200&v=4", "PaperMC"],
        ["DecompMe.png", "decomp.me"],
        ["https://appwrite.io/assets/logomark/logo.png", "Appwrite"],
        ["https://raw.githubusercontent.com/emilk/egui/refs/heads/main/web_demo/favicon.ico", "egui"]
    ]
</script>

<div class="flex flex-col items-center">
    
    <div class="w-full lg:flex pb-10 lg:pb-30">
        <img src="banner.png" alt="banner" class="w-full lg:w-[45%] object-left object-contain" style="image-rendering: pixelated">
        <div class="w-full lg:w-[50%] flex-col lg:pt-20">
            <h1 class="text-9xl text-center w-full">hai, i'm</h1>
            <img src="name_plate.png" alt="nibbles" class="h-[15em] object-center w-full object-contain" style="image-rendering: pixelated">
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
</div>