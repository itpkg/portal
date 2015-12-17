import React from 'react'
import {Alert} from 'react-bootstrap';

var parse = require('url-parse');



const message = React.createClass({
    render(){
        var query = parse(window.location.href, true).query;

        return (<div className="row">
                <div className="col-md-offset-1 col-md-10">
                    <Alert bsStyle={query.ok =='true' ? "success":"danger"}>
                        <h4>{query.msg}</h4>
                    </Alert>
                </div>
            </div>
        )
    }
});

module.exports = {
    Message: message
};