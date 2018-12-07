-- Changes all name of AA@nyu.edu's family group to "fam"
-- Updates all references accordingly
INSERT INTO Friend_Group
    (fg_name, owner_email, description)
VALUES
    ("fam", "AA@nyu.edu", "Ann's Family");

UPDATE Belong
SET fg_name="fam"
WHERE fg_name="family"
AND owner_email="AA@nyu.edu";

UPDATE Share
SET fg_name="fam"
WHERE fg_name="family"
AND owner_email="AA@nyu.edu";

DELETE FROM Friend_Group
WHERE fg_name="family"
AND owner_email="AA@nyu.edu";