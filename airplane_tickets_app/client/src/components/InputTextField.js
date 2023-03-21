import { TextField } from '@mui/material';
import { Controller } from 'react-hook-form';
import React from "react";

export default function InputTextField(props) {
return (
    <Controller
        name={props.name}
        control={props.control}
        defaultValue=""
        render={({ field: { onChange, value }, fieldState: { error } }) => (
          <TextField
            label={props.label}
            type={props.type}
            variant="filled"
            value={value}
            fullWidth
            onChange={onChange}
            error={!!error}
            helperText={error ? error.message : null}
          />
        )}
        rules={props.rules}
    />
);
}