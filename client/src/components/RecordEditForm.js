import { Button, Card, Input, InputAdornment, MenuItem, Select, TextField } from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import React, { useState } from 'react';

const useStyles = makeStyles(theme => ({
    editFormContainer: {
        maxWidth: '300px',
        padding: '40px',
    },
    editForm: {
        display: 'flex',
        flexDirection: 'column',
    },
    editFormItem: {
        textAlign: 'left',
        marginBottom: '15px',
    },
}))

const RecordEditForm = ({ record, saveRecord }) => {
    const classes = useStyles()
    if (record === undefined) {
        record = {}
    }
    const [name, setName] = useState(record.name)
    const [type, setType] = useState(record.type)
    const [balance, setBalance] = useState(record.balance)

    const onSubmit = async (e) => {
        e.preventDefault()
        await saveRecord({
            id: record.id,
            name: name,
            type: type,
            balance: balance,
        })
    }



    return (
        <Card className={classes.editFormContainer}>
            <form onSubmit={onSubmit} className={classes.editForm}>
                <TextField
                    autoFocus
                    required
                    inputProps={{
                        maxLength: 30,
                    }}
                    label="Name"
                    className={classes.editFormItem}
                    value={name}
                    onChange={e => { setName(e.target.value) }}
                    margin="normal"
                />
                <Select
                    labelId="record-type"
                    id="record-type"
                    value={type}
                    defaultValue='Asset'
                    placeholder="Type"
                    className={classes.editFormItem}
                    onChange={e => { setType(e.target.value) }}
                >
                    <MenuItem value={"Asset"}>Asset</MenuItem>
                    <MenuItem value={"Liability"}>Liability</MenuItem>
                </Select>
                <Input
                    className={classes.editFormItem}
                    id="record-balance"
                    value={balance}
                    onChange={e => { setBalance(e.target.value) }}
                    placeholder="Balance"
                    startAdornment={<InputAdornment position="start">$</InputAdornment>}
                />
                <Button variant='contained' color='primary' type='submit' className={classes.formAction}>Save</Button>
            </form>
        </Card>
    )
}

export default RecordEditForm;
