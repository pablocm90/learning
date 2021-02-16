class HotChocolate
  def drink
    'Delicious'
  end

  def discard
    'Plop!'
  end

  def purchase(number)
    "Awesome, I just puchased #{number} chocolate beverage !"
  end
end

RSpec.describe HotChocolate do 
  it 'confirms that an object can respond to a method' do 
    expect(subject).to respond_to(:drink, :discard, :purchase)
  end

  it 'confirms that an object can respond to a method with arguments' do 
    expect(subject).to respond_to(:purchase).with(1).arguments
  end
end
