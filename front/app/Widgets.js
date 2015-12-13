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

const aboutUs = React.createClass({
    render(){
        return (<div>About us</div>)
    }
});

const home = React.createClass({
    render(){
        return (<div>Home</div>)
    }
});

const noMatch = React.createClass({
    render(){
        return (<div>No match</div>)
    }
});

module.exports = {
    Home: home,
    AboutUs: aboutUs,
    NoMatch: noMatch,
    Version: 'v20151212'
};