import React from 'react'
import {Link} from 'react-router'

const header = React.createClass({
    render(){
        return (<div> header </div>)
    }
});

const footer = React.createClass({
    render(){
        return (<div>footer</div>)
    }
});

const noMatch = React.createClass({
    render(){
        return (<div>No match</div>)
    }
});


const version = 'v20151212';

module.exports = {
    Header: header,
    Footer: footer,
    NoMatch: noMatch,
    Version: version
};