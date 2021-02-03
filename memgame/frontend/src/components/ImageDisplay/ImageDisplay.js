import { Component } from "react";

const slides = Array(6).fill().map((item, i) => ( <
    img key = { i }
    src = { require(`../Images${i + 1}.png`) }
    className = "slide" / >
));

class ImageDisplay extends Component {
    render() {
        return (slides)
    }

}

export default ImageDisplay