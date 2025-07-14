-- This function creates a single track with a given offset.
-- It's a template for the repeating track data.
local function create_track(offset)
	return {
		interpolator = "linear",
		value_type = "hsv",
		value_1 = {
			{ x = 0, y = 0.68 },
			{ x = 0.25, y = 0.68 },
			{ x = 0.5, y = 0.71 },
		},
		value_2 = {
			{ x = 0, y = 1 },
			{ x = 1, y = 1 },
		},
		value_3 = {
			{ x = 0, y = 0 },
			{ x = 0.2, y = 1 },
			{ x = 0.25, y = 1 },
			{ x = 0.5, y = 0 },
		},
		brightness = {
			{ x = 0, y = 1.0 },
			{ x = 1, y = 1.0 },
		},
		offset = offset,
	}
end

-- The main animation table to be returned
local animation = {
	name = "Animation 1",
	fps = 60,
	duration_seconds = 1,
	tracks = {}, -- An empty table to hold the tracks
}

-- Use a loop to generate the 17 tracks with varying offsets.
-- This is much cleaner than defining all 17 by hand.
for i = 0, 17 do
	local current_offset = -i / 18 / 2
	table.insert(animation.tracks, create_track(current_offset))
end

-- Return the final, complete animation table
return animation
