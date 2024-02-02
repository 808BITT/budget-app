import {createSignal, For} from "solid-js";

import './ImageGrid.css';


function ImageGrid() {
    const [images, setImages] = createSignal([]);

    // on load fetch images from picsum
    window.addEventListener('load', () => {
        fetch('https://picsum.photos/v2/list')
            .then((response) => response.json())
            .then((data) => {
                data.forEach((item) => item.id++);
                setImages(data.map((item) => item.id));
            });
    });


    function handleClick() {
        console.log('Image clicked');
    }

    return (
        <div class="ImageGridContainer">
            <For each={images()}>
                {(item) =>
                    <a class="ImageGridItem" onClick={handleClick}>
                        <img src={`https://picsum.photos/seed/${item}/800`} alt={`Image ${item}`} />
                    </a>
                }
            </For>
        </div>
    );
}

export default ImageGrid;