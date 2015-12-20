import React from 'react';
import {Alert} from 'react-bootstrap'
import parse from 'url-parse'

export const Home = React.createClass({
    render(){ //todo
        return (<div>
            home
        </div>)
    }
});


export const AboutUs = React.createClass({
    render(){ //todo
        return (<div>
            about us
        </div>)
    }
});


export const NoMatch = React.createClass({
    render(){
        return (<div className="col-md-offset-1 col-md-10">
            <br/>
            <Alert bsStyle="danger">
                <h4>{i18next.t("no_match")}</h4>
            </Alert>
        </div>)
    }
});


export const Message = React.createClass({
    getInitialState: function () {
        var query = parse(window.location.href, true).query;
        return {ok: query.ok == 'true', message: query.msg}
    },
    render(){
        return (<div className="row">
                <div className="col-md-offset-1 col-md-10">
                    <Alert bsStyle={ this.state.ok? "success":"danger"}>
                        <h4>{this.state.message}</h4>
                    </Alert>
                </div>
            </div>
        )
    }
});