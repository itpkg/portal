import React from 'react'

const Header = React.createClass({
    render(){
        return (<div> header </div>)
    }
});

const Footer = React.createClass({
    render(){
        return (<div>footer</div>)
    }
});

module.exports = {
    App: React.createClass({
        render(){
            return (
                <div>
                    App
                    <br/>
                    {this.props.children}
                </div>)
        }
    }),
    HomePage: React.createClass({
        render(){
            return (<div>Home</div>)
        }
    }),
    AboutUsPage: React.createClass({
        render(){
            return (<div>About us</div>)
        }
    }),
    NoMatchPage: React.createClass({
        render(){
            return (<div>No match</div>)
        }
    }),
    Version: 'v20151212'
};