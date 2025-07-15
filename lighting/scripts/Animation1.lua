-- Seed the random number generator, usually done once at the start of your program
math.randomseed(os.time())

-- This function creates a single track with a given offset.
-- It's a template for the repeating track data.
local function create_track(offset)
	local e = math.pow(-offset * 0.1 + 0.1, 1.5)
	return {
		interpolator = "linear",
		value_type = "hsv",
		value_1 = {
			{ x = 0, y = 0.60 + e },
			{ x = 0.5, y = 0.67 + e },
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

local function shuffle_in_place(tbl)
	for i = #tbl, 2, -1 do
		local j = math.random(i)
		tbl[i], tbl[j] = tbl[j], tbl[i]
	end
end

local function reverse_in_place(tbl)
	local n = #tbl
	for i = 1, n / 2 do
		tbl[i], tbl[n - i + 1] = tbl[n - i + 1], tbl[i]
	end
end

-- The main animation table to be returned
local animation = {
	name = "Animation 1",
	fps = 60,
	duration_seconds = 5,
	tracks = {}, -- An empty table to hold the tracks
}

-- Use a loop to generate the 17 tracks with varying offsets.
-- This is much cleaner than defining all 17 by hand.
for i = 0, 17 do
	local current_offset = -i / 18 / 2
	table.insert(animation.tracks, create_track(current_offset))
	reverse_in_place(animation.tracks)
end

-- Return the final, complete animation table
return animation
