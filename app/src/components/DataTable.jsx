import { DataTable} from 'primereact/datatable';
import { Column } from 'primereact/column';
import 'primereact/resources/themes/lara-light-indigo/theme.css';
import 'primereact/resources/primereact.min.css';
import React, { useState, useRef } from 'react';
import 'primereact/resources/themes/lara-light-indigo/theme.css';
import { FilterMatchMode, PrimeIcons } from 'primereact/api';
import { Messages } from 'primereact/messages';
import { Button } from 'primereact/button';
import { ButtonGroup } from 'primereact/buttongroup';
import { Toast } from 'primereact/toast';
import 'primeicons/primeicons.css';
import { ConfirmDialog, confirmDialog } from 'primereact/confirmdialog';

const DataTableComponent = () => {
    const [oFilters, setFilters] = useState({
        global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    });
    const aMessages = useRef(null);

    const addMessage = () => {
        aMessages.current.show([
            { severity: 'info', life: 3172, summary: 'Info', detail: 'Here You got item with n Value, w Weight and n/w Ratio ;)', sticky: false, closable: false }
        ]);
    }

    const [aData, setData] = useState([
        {
            id : 1,
            value : 25,
            weight : 50
        },
        {
            id : 2,
            value : 69,
            weight : 12
        },
        {
            id : 3,
            value : 22,
            weight : 99
        },
        {
            id : 4,
            value : 23,
            weight : 69
        }
    ]);

    const fnAddEmptyRow = () => {
        const oNewRow = {
            id : aData.length + 1,
            value : null,
            weight : ''
        };
        setData([...aData, oNewRow]);
    };


    const fnRemoveRow = (oRowData) => {
        confirmDialog({
            message: 'Are You sure ??? \n Do You want to delete this record?',
            header: 'Confirm',
            icon: 'pi pi-exclamation-triangle',
            accept: () => {
                setData(aData.filter(oRow => oRow.id !== oRowData.id));
                fnShowSuccess();
            },
            reject: () => {fnShowErrMessage()}
        });
    };

    const fnEsditRow = (oRowData) => {
        const oReplacedRow = {
            value : null,
            weight : ''
        };
        setData([...aData, oReplacedRow]);
    };

    const fnEsditRowComplete = (oRowData) => {
        const oReplacedRow = {
            value : null,
            weight : ''
        };
        setData([...aData, oReplacedRow]);
    };

    const oToast = useRef(null);
    const oMessage = useRef(null);

    const fnShowSuccess = () => {
        oToast.current.show({severity: 'success', summary: 'Success Message', detail: 'Message Content', life: 3000});
    };

    const fnShowError = () => {
        oToast.current.show({ severity: 'error', summary: 'Error Message', detail: 'Message Content', life: 3000 });
    };

    const fnShowErrMessage = () => {
        oMessage.current.show( {sticky: true, severity: 'info', summary: 'Info', detail: 'You cancelled delete row', closable: false});
    }
    

    const actionBodyTemplate = (oRowData) => {
        return (
            <React.Fragment>
                <ButtonGroup>
                    <Button severity="danger" className='button-margin' icon={PrimeIcons.MINUS} onClick={() => fnRemoveRow(oRowData)}/>
                    <Button severity="info" size="small" rounded className='button-margin' icon={PrimeIcons.INFO} onClick={addMessage}/>
                    <Button severity="warning" className='button-margin' icon={PrimeIcons.USER_EDIT} onClick={() => fnEsditRow(oRowData)}/>
                        <Messages ref={aMessages}/>
                </ButtonGroup>
            </React.Fragment>
        );
    };

    return (
        <div className="App">
            <Messages ref={oMessage}/>
            <Toast ref={oToast}/>
            <ConfirmDialog/>
            <div style={{ display: "flex" }}>
            <Button severity="info" style={{ marginLeft: "auto" }} rounded  className="button-margin" icon={PrimeIcons.PLUS} onClick={fnAddEmptyRow}/>
            </div>
            <DataTable
                id="knapsackTable"
                editMode='cell'
                value={aData}
                scrollable
                dataKey="id"
                selectionMode="single"
                sortField="name"
                sortOrder={-1}
                sortMode="multiple"
                filters={oFilters}
                onRowEditChange={fnEsditRow}
                onRowEditComplete={fnEsditRowComplete}
                paginator
                rows={aData.length}
                rowsPerPageOptions={[25, 50, 100]}
                totalRecords={aData.length}
                tableStyle={{ minWidth: '60rem' }}
                showGridlines
                stripedRows
                scrollHeight='500px'>
                <Column field="id" header="ID" hidden/>
                <Column field="value" header="Value [$]" sortable/>
                <Column field="weight" header="Weight [Kg]" sortable/>
                <Column body={actionBodyTemplate} expander header="Actions"/>
            </DataTable>
        </div>
    );
}

export default DataTableComponent;