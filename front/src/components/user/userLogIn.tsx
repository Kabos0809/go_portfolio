import React, { Component } from "react"
import { Container, FormControl, FormLabel, TextField, Button, Box } from "@material-ui/core"
import { Alert } from '@material-ui/lab'
import { WithRouter } from 'react-router-dom'
import { User }  from '../../api/user'

const clickLogin = async () => {
    try {
        await User.Login(state.username, state.password)
        props.history.push({pathname: '/admin'})
    } catch (e) {
        this.setState({ errMessage: 'Wrong name or password'})
        alert(this.state)
    }
}