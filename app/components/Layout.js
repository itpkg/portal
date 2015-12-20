import React from 'react';
import {IndexLink, Link} from 'react-router';
import {Header, Footer} from './Widgets'

export const Application = React.createClass({
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
