class Influencer:
    def __init__(self, num_selfies:int, num_bio_link:int):
        self.num_selfies = num_selfies
        self.num_bio_link = num_bio_link
    
    def __repr__(self):
        return f"Influencer({self.num_selfies}, {self.num_bio_link})"    
        

walon = Influencer(100, 10)

def vanity(influencer:Influencer) -> int:
    return (influencer.num_bio_link * 5) + influencer.num_selfies
    
def vanity_sort(group:list[Influencer]) -> list:
    collection:dict[Influencer, int] = {}
    for influencer in group:
        collection[influencer] = vanity(influencer)
    
    list = sorted(collection, key=lambda inf: collection[inf])
    return list
            

theprimeagen = Influencer(100, 1)
pokimane = Influencer(800, 2)
spambot = Influencer(0, 200)
lane = Influencer(10, 2)
badcop = Influencer(1, 2)

list = vanity_sort([spambot, theprimeagen])
print(list)