import React from 'react'
import {Alert} from 'react-bootstrap';

var parse = require('url-parse');


const message = React.createClass({
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

module.exports = {
    Message: message
};