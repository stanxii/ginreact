import React, { Component } from 'react';

class Header extends Component {
    render() {
        return (
            <nav className="navbar navbar-expand-lg navbar-light bg-dark">
                <a style={{ color: '#fff' }} href="/">GoReactRest</a>
            </nav>
        );
    }
}

export default Header;