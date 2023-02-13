-- Function to convert RGB values to a hexadecimal string
local function rgbToHex(r, g, b)
    return string.format("#%02x%02x%02x", r, g, b)
end

-- Function to convert a hexadecimal string to RGB values
local function hexToRgb(hex)
    local r, g, b = hex:match("#(%x%x)(%x%x)(%x%x)")
    return tonumber(r, 16), tonumber(g, 16), tonumber(b, 16)
end

-- Table of color names and their RGB or hexadecimal values
local color_map = {
    ["red"] = {255, 0, 0},
    ["green"] = {0, 255, 0},
    ["blue"] = {0, 0, 255},
    ["yellow"] = "#FFFF00",
    ["purple"] = "#800080",
    ["cyan"] = "#00FFFF",
}

-- Function to return the categorized name of a color
local function getColorName(color)
    -- Check if the input is an RGB triplet
    if type(color) == "table" then
        color = rgbToHex(unpack(color))
    end

    -- Loop through the color map to find a match
    for name, value in pairs(color_map) do
        if color == value or color == rgbToHex(unpack(value)) then
            return name
        end
    end

    -- Return "unknown" if no match was found
    return "unknown"
end

-- Example usage
print(getColorName({255, 0, 0})) -- outputs "red"
print(getColorName("#FF0000")) -- outputs "red"
