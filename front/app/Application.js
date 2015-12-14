import React from 'react';
import {IndexLink} from 'react-router';


const application = React.createClass({
    render(){
        return (
            <div>
                Application
                <br/>
                <IndexLink to="/">Home</IndexLink>
                &nbsp;
                <IndexLink to="/about_us">About Us</IndexLink>
                <br/>
                {this.props.children}
            </div>)
    }
});

const home = React.createClass({
    render(){
        return (<div>
            home
            <br/>
            <IndexLink to="/about_us">About Us</IndexLink>

        </div>)
    }
});

const aboutUs = React.createClass({
    render(){
        return (<div>
            about us
            <br/>
        </div>)
    }
});

const notice = React.createClass({
    render() {
        return <h3>Message {this.props.params.id}</h3>
    }
});

module.exports = {
    Application: application,
    Home: home,
    AboutUs: aboutUs,
    Notice: notice
};