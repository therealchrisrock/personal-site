const scriptTag = document.currentScript as HTMLScriptElement;
const libraryType = scriptTag?.dataset.library; // Read the "data-library" attribute
async function loadLibrary() {
    switch (libraryType) {
        case "easter-egg":
            const EASTER_EGG = await import('./easter-egg');
            await EASTER_EGG.init()
            break
        case "d3":
            const D3_MODULE = await import("./mantracker"); // Lazy-load D3 visualization
            D3_MODULE.initD3Visualization();
            break;
        default:
            console.error("Unknown library type");
    }
}

document.addEventListener('DOMContentLoaded', loadLibrary)
