RSpec.describe 'raise error matcher' do
  def some_method
    x
  end

  it 'can check for any error' do
    expect { some_method }.to raise_error
  end

  it 'can check for specific error' do
    expect { some_method }.to raise_error(NameError)
  end
end
