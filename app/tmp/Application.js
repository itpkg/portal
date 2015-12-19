import React from 'react';
import {IndexLink} from 'react-router';

import {Header, Footer} from './Widgets';

const application = React.createClass({
    render(){
        return (
            <div>
                <Header/>
                <div className="container-fluid">
                    <div className="row">
                        {this.props.children}
                    </div>
                    <hr/>
                    <div>
                        <Footer/>
                    </div>
                </div>
            </div>
        )
    }
});

const home = React.createClass({
    render(){ //todo
        return (<div>
            home
            <br/>
            <IndexLink to="/about_us">About Us</IndexLink>

        </div>)
    }
});

const aboutUs = React.createClass({
    render(){ // todo
        return (<div>
            about us
            <br/>
        </div>)
    }
});

const notice = React.createClass({
    render() { //todo
        return <h3>Message {this.props.params.id}</h3>
    }
});

module.exports = {
    Application: application,
    Home: home,
    AboutUs: aboutUs,
    Notice: notice
};