import { makeStyles } from '@material-ui/styles';
import React from 'react';
import { connect } from 'react-redux';
import RecordTable from '../components/RecordTable';

const useStyles = makeStyles(theme => ({

}))

class RecordTableContainer extends React.Component {
    constructor(props) {
        super(props);
    }

    componentDidMount() {
        console.log("Component did mount")
        this.props.fetchRecords()
    }

    render() {
        console.log('Render records', this.props.records)
        if (this.props.records.records && this.props.records.records.length > 0) {
            console.log(this.props.records.records[0].id)

        }
        return (
            <>
                <RecordTable records={this.props.records} isLoading={this.props.isLoading} error={this.props.error} updateRecord={this.props.updateRecord} />
            </>
        )
    }
}

const mapStateToProps = state => ({
    records: state.record.records,
    isLoading: state.record.isLoading,
    error: state.record.error,
})

const mapDispatchToProps = ({ record: { fetchRecords, updateRecord } }) => ({
    fetchRecords: () => fetchRecords(),
    updateRecord: (id, name, type, balance) => updateRecord(id, name, type, balance),
})

export default connect(mapStateToProps, mapDispatchToProps)(RecordTableContainer);

