import { CircularProgress, Typography } from "@material-ui/core";
import { makeStyles } from '@material-ui/styles';
import React from 'react';


const useStyles = makeStyles(theme => ({
    root: {
        margin: '40px 0',
    },
    positive: {
        color: 'green',
    },
    negative: {
        color: 'red',
    }
}))

export default function NetWorth({ isLoading, netWorth }) {
    const classes = useStyles()
    if (isLoading) {
        return <CircularProgress />
    }

    if (!netWorth) {
        netWorth = 0
    }

    return (
        <div className={classes.root}>
            <Typography variant='h2' >
                Net Worth: <span className={netWorth >= 0 ? classes.positive : classes.negative}>${netWorth.toFixed(2)} </span>
            </Typography>
        </div>
    )
}