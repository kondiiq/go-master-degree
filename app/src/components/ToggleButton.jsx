import { ToggleButton } from "primereact/togglebutton";
import React, { useState } from "react";

function ToggleButtonComponent() {
    const [isChecked, setChecked] = useState(false);

    return (
        <div className="card">
            <ToggleButton size="small"  type="checkbox" onLabel="Dark" offLabel="White" onIcon="pi pi-check" offIcon="pi pi-times" 
                checked={isChecked}  className="w-9rem" />
        </div>
    );
}   

export default ToggleButtonComponent;