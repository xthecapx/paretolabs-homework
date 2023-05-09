import ListItem from "@mui/material/ListItem";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemText from "@mui/material/ListItemText";
import { forwardRef } from "react";

export const Item = forwardRef<HTMLLIElement, { text: string }>(
  ({ text }, ref) => {
    return ref ? (
      <ListItem ref={ref} disablePadding>
        <ListItemButton>
          <ListItemText primary={text} />
        </ListItemButton>
      </ListItem>
    ) : (
      <ListItem disablePadding>
        <ListItemButton>
          <ListItemText primary={text} />
        </ListItemButton>
      </ListItem>
    );
  }
);
