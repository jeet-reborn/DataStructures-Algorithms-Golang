import React, { Component } from "react"

class Header extends Component {
    componentDidMount() {
        document.title = 'Memory Game';
    }

    render() {
        return (
            <div className="header">
                Realtime Memory Game
            </div>
        );
    }
}

export default Header