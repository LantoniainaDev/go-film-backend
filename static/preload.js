const loader = document.querySelector(".loader");
console.log("loader", loader);

window.addEventListener("load",()=>{
    console.log("loader sera supprimee");
    loader.animate([
        {opacity:1},
        {opacity:0},
    ],{duration:500, easing:"ease-in"}).finished.then(()=>loader.style.display = "none"); 
})