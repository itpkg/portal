import React from 'react';
import {Link} from 'react-router';

const application = React.createClass({
    render(){
        return (
            <div>
                Application
                <br/>
                <Link to="/">Home</Link>
                &nbsp;
                <Link to="/about_us">About Us</Link>
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
            <Link to="/about_us">About Us</Link>

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